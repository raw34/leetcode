<?php

namespace book\design_pattern\command;

interface Command
{
    public function execute(): string;
    public function undo(): string;
    public function redo(): string;
}
