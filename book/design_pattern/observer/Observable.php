<?php

namespace book\design_pattern\observer;

interface Observable
{
    public function attach(Observer $observer): void;
    public function detach(Observer $observer): void;
    public function notify(JobPost $jobPost): void;
}