import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import StudentsPage from './pages/StudentsPage';
import SubjectsPage from './pages/SubjectsPage';
import GradesPage from './pages/GradesPage';
import './App.css';

function App() {
  return (
    <Router>
      <nav>
        <Link to="/">Inicio</Link> | 
        <Link to="/students">Estudiantes</Link> | 
        <Link to="/subjects">Materias</Link> | 
        <Link to="/grades">Calificaciones</Link>
      </nav>
      <div className="container">
        <Routes>
          <Route path="/" element={<h1>Bienvenido al Sistema de Control Escolar</h1>} />
          <Route path="/students" element={<StudentsPage />} />
          <Route path="/subjects" element={<SubjectsPage />} />
          <Route path="/grades" element={<GradesPage />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
