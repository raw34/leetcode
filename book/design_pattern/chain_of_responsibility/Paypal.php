<?php

namespace book\design_pattern\chain_of_responsibility;

class Paypal extends Account
{
    protected float $balance;

    public function __construct(float $balance)
    {
        $this->balance = $balance;
    }
}