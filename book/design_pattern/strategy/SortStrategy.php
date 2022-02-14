<?php

namespace book\design_pattern\strategy;

interface SortStrategy
{
    public function sort(array $dataset): string;
}
