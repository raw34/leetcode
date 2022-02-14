<?php

namespace book\design_pattern\state;

class DefaultText implements WritingState
{
    public function write(string $words): string
    {
        return $words;
    }
}