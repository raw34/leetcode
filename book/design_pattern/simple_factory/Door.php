<?php

namespace book\design_pattern\simple_factory;

interface Door
{
    public function getWidth(): float;
    public function getHeight(): float;
}