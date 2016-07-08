package matching
import(
  "testing"
)

func Test(t *testing.T){
  expected := "Hello, World"
  actual := Hello()
  if(expected != actual){
    t.Errorf("expected: %s did not match actual: %s", expected, actual)
  }
}
