<?php

namespace book\design_pattern\visitor;

interface AnimalOperation
{
    public function visitMonkey(Monkey $monkey): string;

    public function visitLion(Lion $lion): string;

    public function visitDolphin(Dolphin $dolphin): string;
}