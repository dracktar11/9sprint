package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	empty := generateRandomElements(0)
	if len(empty) != 0 {
		t.Errorf("для размера 0 ожидали пустой слайс, получили длину %d", len(empty))
	}

	got := generateRandomElements(10)
	if len(got) != 10 {
		t.Errorf("ожидали длину 10, получили %d", len(got))
	}
}

func TestMaximum(t *testing.T) {
	if maximum([]int{}) != 0 {
		t.Error("для пустого слайса ожидали 0")
	}
	if maximum([]int{5}) != 5 {
		t.Error("для одного элемента ожидали 5")
	}
	if maximum([]int{1, 7, 3}) != 7 {
		t.Error("ожидали максимум 7")
	}
}
