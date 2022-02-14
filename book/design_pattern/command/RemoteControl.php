<?php

namespace book\design_pattern\command;

class RemoteControl
{
    public function submit(Command $command): string
    {
        return $command->execute();
    }
}