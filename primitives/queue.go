package main
 
import "fmt"
 
type stack []interface{}
 
func (k *stack) push(s interface{}) {
    *k = append(*k, s)
}
 
func (k *stack) pop() (s interface{}, ok bool) {
    if k.empty() {
        return
    }
    //last := len(*k) - 1
    s = (*k)[0]
    *k = (*k)[1:]
    return s, true
}
 
func (k *stack) peek() (s interface{}, ok bool) {
    if k.empty() {
        return
    }
    last := len(*k) - 1
    s = (*k)[last]
    return s, true
}
 
func (k *stack) empty() bool {
    return len(*k) == 0
}

func main() {
    s := stack{}
    s.push("hi")
    s.push("there")
    
    for !s.empty() {
        l, ok := s.pop()
        if ok {
            fmt.Println(l)
        }
    }
}