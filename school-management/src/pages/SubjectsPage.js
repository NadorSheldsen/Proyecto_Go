import React, { useState, useEffect } from 'react';
import { getAllSubjects, createSubject, updateSubject, deleteSubject } from '../services/subjectsService';
import SubjectForm from '../components/SubjectForm';
import Modal from '../components/Modal';
import './SubjectsPage.css';

function SubjectsPage() {
  const [subjects, setSubjects] = useState([]);
  const [selectedSubject, setSelectedSubject] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    fetchSubjects();
  }, []);

  const fetchSubjects = async () => {
    try {
      const response = await getAllSubjects();
      setSubjects(response.data);
    } catch (error) {
      console.error("Error al cargar materias:", error);
    }
  };

  const openCreateSubjectModal = () => {
    setSelectedSubject(null); 
    setIsModalOpen(true);
  };

  const openEditSubjectModal = (subject) => {
    setSelectedSubject(subject); 
    setIsModalOpen(true);
  };

  const handleSaveSubject = async (data) => {
    try {
      if (selectedSubject) {
        await updateSubject(selectedSubject.id, data); 
      } else {
        await createSubject(data);
      }
      fetchSubjects(); 
      setIsModalOpen(false);
    } catch (error) {
      console.error("Error al guardar materia:", error);
    }
  };

  const handleDeleteSubject = async (id) => {
    try {
      await deleteSubject(id);
      fetchSubjects(); 
    } catch (error) {
      console.error("Error al eliminar materia:", error);
    }
  };

  return (
    <div className="container">
      <h1>Materias</h1>
      <button onClick={openCreateSubjectModal}>Agregar Materia</button>
      <table>
        <thead>
          <tr>
            <th>Nombre</th>
            <th>Código</th>
            <th>Créditos</th>
            <th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          {subjects.map((subject) => (
            <tr key={subject.id}>
              <td>{subject.name}</td>
              <td>{subject.code}</td>
              <td>{subject.credits}</td>
              <td className="actions">
                <button onClick={() => openEditSubjectModal(subject)}>Editar</button>
                <button onClick={() => handleDeleteSubject(subject.id)}>Eliminar</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      {isModalOpen && (
        <Modal onClose={() => setIsModalOpen(false)}>
          <div className="modal-content">
            <SubjectForm
              subject={selectedSubject}
              onSave={handleSaveSubject}
            />
          </div>
        </Modal>
      )}
    </div>
  );
}

export default SubjectsPage;
