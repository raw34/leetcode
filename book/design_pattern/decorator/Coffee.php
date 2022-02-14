<?php

namespace book\design_pattern\decorator;

interface Coffee
{
    public function getCost(): float;
    public function getDescription(): string;
}