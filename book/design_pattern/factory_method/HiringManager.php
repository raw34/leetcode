<?php

namespace book\design_pattern\factory_method;

abstract class HiringManager
{
    // Factory method
    abstract protected function makeInterviewer(): Interviewer;

    public function takeInterview(): string
    {
        return $this->makeInterviewer()->askQuestions();
    }
}