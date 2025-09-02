package main

import (
    "fmt"
)

func main() {
    // Create slice with 0 length but capacity of 3
    sli := make([]int, 0, 3)

    // Get the address of the underlying array
    // We'll track if this changes after append
    fmt.Printf("Initial: len=%d cap=%d\n", len(sli), cap(sli))

    // Store original capacity to track if underlying array changes
    originalCap := cap(sli)

    // Append 1st element (should use existing capacity)
    sli = append(sli, 100)
    fmt.Printf("After 1st append: len=%d cap=%d value=%v\n", len(sli), cap(sli), sli)
    if cap(sli) == originalCap {
        fmt.Println("  ✓ Same underlying array (capacity unchanged)")
    }

    // Append 2nd element (should still use existing capacity)
    sli = append(sli, 200)
    fmt.Printf("After 2nd append: len=%d cap=%d value=%v\n", len(sli), cap(sli), sli)
    if cap(sli) == originalCap {
        fmt.Println("  ✓ Same underlying array (capacity unchanged)")
    }

    // Append 3rd element (fills to capacity)
    sli = append(sli, 300)
    fmt.Printf("After 3rd append: len=%d cap=%d value=%v\n", len(sli), cap(sli), sli)
    if cap(sli) == originalCap {
        fmt.Println("  ✓ Same underlying array (capacity unchanged)")
    }

    // Append 4th element (exceeds capacity - should trigger reallocation)
    sli = append(sli, 400)
    fmt.Printf("After 4th append: len=%d cap=%d value=%v\n", len(sli), cap(sli), sli)
    if cap(sli) != originalCap {
        fmt.Printf("  ✗ NEW underlying array allocated! (capacity changed from %d to %d)\n", originalCap, cap(sli))
    }

    fmt.Println("\n--- Demonstration with pointer comparison ---")

    // More definitive proof using pointers
    sli2 := make([]int, 0, 3)

    // Trick: extend to cap to get pointer to underlying array
    extended := sli2[:cap(sli2)]
    ptr1 := &extended[0]
    fmt.Printf("Pointer to underlying array: %p\n", ptr1)

    // Append within capacity
    sli2 = append(sli2, 111)
    extended = sli2[:cap(sli2)]
    ptr2 := &extended[0]
    fmt.Printf("After append (within cap):   %p", ptr2)
    if ptr1 == ptr2 {
        fmt.Println(" ← Same address!")
    }

    // Append beyond capacity
    sli2 = append(sli2, 222, 333, 444)
    extended = sli2[:cap(sli2)]
    ptr3 := &extended[0]
    fmt.Printf("After append (beyond cap):   %p", ptr3)
    if ptr1 != ptr3 {
        fmt.Println(" ← Different address!")
    }
}
