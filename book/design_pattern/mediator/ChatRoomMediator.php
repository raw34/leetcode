<?php

namespace book\design_pattern\mediator;

interface ChatRoomMediator
{
    public function showMessage(User $user, string $message): string;
}