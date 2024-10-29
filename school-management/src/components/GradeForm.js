import React, { useState, useEffect } from 'react';

function GradeForm({ grade, onSave }) {
  const [formData, setFormData] = useState({
    studentId: grade ? grade.studentId : '',
    subjectId: grade ? grade.subjectId : '',
    grade: grade ? grade.grade : ''
  });

  useEffect(() => {
    if (grade) {
      setFormData({
        studentId: grade.studentId,
        subjectId: grade.subjectId,
        grade: grade.grade
      });
    }
  }, [grade]);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSave(formData);
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>ID del Estudiante:</label>
      <input
        type="text"
        name="studentId"
        value={formData.studentId}
        onChange={handleChange}
        required
        placeholder="ID del estudiante"
      />

      <label>ID de la Materia:</label>
      <input
        type="text"
        name="subjectId"
        value={formData.subjectId}
        onChange={handleChange}
        required
        placeholder="ID de la materia"
      />

      <label>Calificación:</label>
      <input
        type="number"
        name="grade"
        value={formData.grade}
        onChange={handleChange}
        required
        placeholder="Calificación"
        min="0"
        max="100"
      />

      <button type="submit">Guardar</button>
    </form>
  );
}

export default GradeForm;
