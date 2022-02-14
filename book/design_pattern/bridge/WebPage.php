<?php

namespace book\design_pattern\bridge;

interface WebPage
{
    public function __construct(Theme $theme);
    public function getContent(): string;
}