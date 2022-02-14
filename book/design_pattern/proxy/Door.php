<?php

namespace book\design_pattern\proxy;

interface Door
{
    public function open(): string;
    public function close(): string;
}