<?php

namespace book\design_pattern\strategy;

class Sorter
{
    private SortStrategy $sorter;

    public function __construct(SortStrategy $sorter)
    {
        $this->sorter = $sorter;
    }

    public function sort(array $dataset): string
    {
        return $this->sorter->sort($dataset);
    }
}