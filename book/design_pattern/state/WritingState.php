<?php

namespace book\design_pattern\state;

interface WritingState
{
    public function write(string $words): string;
}