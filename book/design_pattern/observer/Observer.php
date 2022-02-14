<?php

namespace book\design_pattern\observer;

interface Observer
{
    public function update(JobPost $jobPost): void;
}