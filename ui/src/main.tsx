import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { HashRouter, Link, Route, Routes } from 'react-router'
import '@mantine/core/styles.css';
import { Button, MantineProvider } from '@mantine/core';
import AppLayout from './layouts/AppLayout';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <MantineProvider>
      <HashRouter>
        <Routes>
          <Route element={<AppLayout />}>
            <Route path="/" element={<div>Hello, world! <Link to="/about">About</Link></div>} />
            <Route path="/about" element={<Button variant="filled">Button</Button>} />
            <Route path="/contact" element={<div>Contact</div>} />
          </Route>
        </Routes>
      </HashRouter>
    </MantineProvider>
  </StrictMode>,
)
