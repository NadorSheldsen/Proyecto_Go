import React, { useState, useEffect } from 'react';
import { getAllGrades, createGrade, updateGrade, deleteGrade } from '../services/gradesService';
import GradeForm from '../components/GradeForm';
import Modal from '../components/Modal';
import './GradesPage.css';

function GradesPage() {
  const [grades, setGrades] = useState([]);
  const [selectedGrade, setSelectedGrade] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    fetchGrades();
  }, []);

  const fetchGrades = async () => {
    try {
      const response = await getAllGrades();
      setGrades(response.data); 
    } catch (error) {
      console.error("Error al cargar calificaciones:", error);
    }
  };

  const openCreateGradeModal = () => {
    setSelectedGrade(null);
    setIsModalOpen(true);
  };

  const openEditGradeModal = (grade) => {
    setSelectedGrade(grade); 
    setIsModalOpen(true);
  };

  const handleSaveGrade = async (data) => {
    try {
      if (selectedGrade) {
        await updateGrade(selectedGrade.id, data);
      } else {
        await createGrade(data);
      }
      fetchGrades(); 
      setIsModalOpen(false);
    } catch (error) {
      console.error("Error al guardar calificación:", error);
    }
  };

  const handleDeleteGrade = async (id) => {
    try {
      await deleteGrade(id);
      fetchGrades(); 
    } catch (error) {
      console.error("Error al eliminar calificación:", error);
    }
  };

  return (
    <div className="container">
      <h1>Calificaciones</h1>
      <button onClick={openCreateGradeModal}>Agregar Calificación</button>
      <table>
        <thead>
          <tr>
        
            <th>Estudiante</th>
            <th>Materia</th>
            <th>Puntaje</th>
            <th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          {grades.map((grade) => (
            <tr key={grade.id}> 
              <td>{grade.studentName}</td> 
              <td>{grade.subjectName}</td> 
              <td>{grade.score}</td> 
              <td className="actions">
                <button onClick={() => openEditGradeModal(grade)}>Editar</button>
                <button onClick={() => handleDeleteGrade(grade.id)}>Eliminar</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      {isModalOpen && (
        <Modal onClose={() => setIsModalOpen(false)}>
          <div className="modal-content">
            <GradeForm
              grade={selectedGrade} 
              onSave={handleSaveGrade}
            />
          </div>
        </Modal>
      )}
    </div>
  );
}

export default GradesPage;
