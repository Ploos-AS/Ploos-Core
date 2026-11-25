package hosting

import "testing"

func TestProjectName(t *testing.T) {
	a, err := ProjectName("org-a", "resource")
	if err != nil {
		t.Fatal(err)
	}
	b, _ := ProjectName("org-a", "resource")
	c, _ := ProjectName("org-b", "resource")
	if a != b || a == c {
		t.Fatal("invalid deployment scope")
	}
	if _, err := ProjectName("", "resource"); err == nil {
		t.Fatal("missing scope accepted")
	}
	if _, err := ProjectName("org\x00a", "resource"); err == nil {
		t.Fatal("invalid scope accepted")
	}
}
