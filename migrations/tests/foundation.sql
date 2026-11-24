BEGIN;
DO $$
DECLARE
 a uuid;
 b uuid;
 p uuid;
 s uuid;
BEGIN
 INSERT INTO organizations(display_name,kind) VALUES ('A','association') RETURNING id INTO a;
 INSERT INTO organizations(display_name,kind) VALUES ('B','business') RETURNING id INTO b;
 INSERT INTO plans(product_code,provider_code,code,seat_limit) VALUES ('cloud','ploos','starter',5) RETURNING id INTO p;
 INSERT INTO subscriptions(organization_id,plan_id) VALUES (a,p) RETURNING id INTO s;
 INSERT INTO entitlements VALUES (a,s,'seats',5);
 BEGIN
  INSERT INTO entitlements VALUES (b,s,'seats',5);
  RAISE EXCEPTION 'cross-tenant entitlement accepted';
 EXCEPTION WHEN foreign_key_violation THEN NULL;
 END;
 BEGIN
  INSERT INTO provisioning_outbox(organization_id,subscription_id,resource_id,correlation_id,idempotency_key,event_type)
  VALUES (b,s,'resource',gen_random_uuid(),'test','create');
  RAISE EXCEPTION 'cross-tenant event accepted';
 EXCEPTION WHEN foreign_key_violation THEN NULL;
 END;
 IF EXISTS (SELECT 1 FROM providers WHERE enabled) THEN
  RAISE EXCEPTION 'provider enabled by default';
 END IF;
END $$;
ROLLBACK;
