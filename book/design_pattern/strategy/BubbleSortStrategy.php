<?php

namespace book\design_pattern\strategy;

class BubbleSortStrategy implements SortStrategy
{
    public function sort(array $dataset): string
    {
        // Do sorting
        return "Sorting using bubble sort";
    }
}