<?php

namespace book\design_pattern\observer;

use PHPUnit\Framework\TestCase;

/*
// Create subscribers
$johnDoe = new JobSeeker('John Doe');
$janeDoe = new JobSeeker('Jane Doe');

// Create publisher and attach subscribers
$jobPostings = new EmploymentAgency();
$jobPostings->attach($johnDoe);
$jobPostings->attach($janeDoe);

// Add a new job and see if subscribers get notified
$jobPostings->addJob(new JobPost('Software Engineer'));

// Output
// Hi John Doe! New job posted: Software Engineer
// Hi Jane Doe! New job posted: Software Engineer
*/

class ObserverTest extends TestCase
{
    public function test01(): void
    {
        // Create subscribers
        $johnDoe = new JobSeeker('John Doe');
        $janeDoe = new JobSeeker('Jane Doe');

        // Create publisher and attach subscribers
        $jobPostings = new EmploymentAgency();
        $jobPostings->attach($johnDoe);
        $jobPostings->attach($janeDoe);

        // Add a new job and see if subscribers get notified
        $jobPostings->addJob(new JobPost('Software Engineer'));

        // Output
        // Hi John Doe! New job posted: Software Engineer
        // Hi Jane Doe! New job posted: Software Engineer
    }
}
