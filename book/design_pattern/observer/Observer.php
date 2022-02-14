<?php

namespace book\design_pattern\observer;

interface Observer
{
    public function onJobPosted(JobPost $job);
}