import React, { useState, useEffect } from 'react';
import { getAllStudents, createStudent, updateStudent, deleteStudent } from '../services/studentsService';
import StudentForm from '../components/StudentForm';
import Modal from '../components/Modal';
import './StudentsPage.css';

function StudentsPage() {
  const [students, setStudents] = useState([]);
  const [selectedStudent, setSelectedStudent] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    fetchStudents();
  }, []);

  const fetchStudents = async () => {
    try {
      const response = await getAllStudents();
      setStudents(response.data);
    } catch (error) {
      console.error("Error al cargar estudiantes:", error);
    }
  };

  const openCreateStudentModal = () => {
    setSelectedStudent(null); // 
    setIsModalOpen(true);
  };

  const openEditStudentModal = (student) => {
    setSelectedStudent(student); 
    setIsModalOpen(true);
  };

  const handleSaveStudent = async (data) => {
    try {
      if (selectedStudent) {
        await updateStudent(selectedStudent.id, data);
      } else {
        await createStudent(data);
      }
      fetchStudents(); 
      setIsModalOpen(false);
    } catch (error) {
      console.error("Error al guardar estudiante:", error);
    }
  };

  const handleDeleteStudent = async (id) => {
    try {
      await deleteStudent(id);
      fetchStudents(); 
    } catch (error) {
      console.error("Error al eliminar estudiante:", error);
    }
  };

  return (
    <div className="container">
      <h1>Estudiantes</h1>
      <button onClick={openCreateStudentModal}>Agregar Estudiante</button>
      <table>
        <thead>
          <tr>
            <th>Nombre</th>
            <th>Grupo</th>
            <th>Email</th>
            <th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          {students.map((student) => (
            <tr key={student.id}>
              <td>{student.name}</td>
              <td>{student.group}</td>
              <td>{student.email}</td>
              <td className="actions">
                <button onClick={() => openEditStudentModal(student)}>Editar</button>
                <button onClick={() => handleDeleteStudent(student.id)}>Eliminar</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      {isModalOpen && (
        <Modal onClose={() => setIsModalOpen(false)}>
          <div className="modal-content">
            <StudentForm
              student={selectedStudent}
              onSave={handleSaveStudent}
            />
          </div>
        </Modal>
      )}
    </div>
  );
}

export default StudentsPage;
