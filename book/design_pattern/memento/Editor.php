<?php

namespace book\design_pattern\memento;

class Editor
{
    private string $content = '';

    public function type(string $words)
    {
        $this->content .= ' ' . $words;
    }

    public function getContent(): string
    {
        return $this->content;
    }

    public function save(): EditorMemento
    {
        return new EditorMemento($this->content);
    }

    public function restore(EditorMemento $memento): void
    {
        $this->content = $memento->getContent();
    }
}