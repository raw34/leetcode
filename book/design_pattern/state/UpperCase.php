<?php

namespace book\design_pattern\state;

class UpperCase implements WritingState
{
    public function write(string $words): string
    {
        return strtoupper($words);
    }
}