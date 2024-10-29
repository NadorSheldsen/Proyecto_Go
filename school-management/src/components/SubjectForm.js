import React, { useState, useEffect } from 'react';

function SubjectForm({ subject, onSave }) {
  const [formData, setFormData] = useState({
    name: subject ? subject.name : '',
    code: subject ? subject.code : '',
    credits: subject ? subject.credits : ''
  });

  useEffect(() => {
    if (subject) {
      setFormData({
        name: subject.name,
        code: subject.code,
        credits: subject.credits
      });
    }
  }, [subject]);

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
      <label>Código:</label>
      <input type="text" name="code" value={formData.code} onChange={handleChange} required />
      <label>Créditos:</label>
      <input type="number" name="credits" value={formData.credits} onChange={handleChange} required />
      <button type="submit">Guardar</button>
    </form>
  );
}

export default SubjectForm;
