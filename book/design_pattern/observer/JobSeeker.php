<?php

namespace book\design_pattern\observer;

class JobSeeker implements Observer
{
    private string $name;

    public function __construct(string $name)
    {
        $this->name = $name;
    }

    public function onJobPosted(JobPost $job)
    {
        // Do something with the job posting
        echo 'Hi ' . $this->name . '! New job posted: ' . $job->getTitle();
    }
}