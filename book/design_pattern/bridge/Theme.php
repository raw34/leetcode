<?php

namespace book\design_pattern\bridge;

interface Theme
{
    public function getColor(): string;
}