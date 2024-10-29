import React, { useState, useEffect } from 'react';

function StudentForm({ student, onSave }) {
  const [formData, setFormData] = useState({
    name: student ? student.name : '',
    group: student ? student.group : '',
    email: student ? student.email : ''
  });

  useEffect(() => {
    if (student) {
      setFormData({
        name: student.name,
        group: student.group,
        email: student.email
      });
    }
  }, [student]);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSave(formData);
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>Nombre:</label>
      <input type="text" name="name" value={formData.name} onChange={handleChange} required />
      <label>Grupo:</label>
      <input type="text" name="group" value={formData.group} onChange={handleChange} required />
      <label>Email:</label>
      <input type="email" name="email" value={formData.email} onChange={handleChange} required />
      <button type="submit">Guardar</button>
    </form>
  );
}

export default StudentForm;
