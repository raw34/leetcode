<?php

namespace book\design_pattern\bridge;

class About implements WebPage
{
    protected $theme;

    public function __construct(Theme $theme)
    {
        $this->theme = $theme;
    }

    public function getContent(): string
    {
        return "About page in " . $this->theme->getColor();
    }
}