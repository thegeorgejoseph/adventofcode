package main


func Two(s string) int {
    cache := make(map[string]bool)
    left, right := 0, 0

    for right < len(s) {
        if _, ok := cache[string(s[right])]; !ok {
           if (right - left + 1) == 14 {
               break
           }
           cache[string(s[right])] = true
           right++
           continue
        }
        delete(cache, string(s[left]))
        left++
    }
    return right + 1
}
