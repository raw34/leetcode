<?php

namespace book\design_pattern\strategy;

class QuickSortStrategy implements SortStrategy
{
    public function sort(array $dataset): string
    {
        // Do sorting
        return "Sorting using quick sort";
    }
}