<?php

namespace book\design_pattern\factory_method;

class CommunityExecutive implements Interviewer
{
    public function askQuestions(): string
    {
        return 'Asking about community building.';
    }
}