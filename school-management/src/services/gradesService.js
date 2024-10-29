// src/services/gradesService.js
import axiosInstance from '../axiosConfig';

export async function getAllGrades() {
  const response = await axiosInstance.get('/grades');
  return response.data;
}

export async function createGrade(gradeData) {
  const response = await axiosInstance.post('/escuela/grades', gradeData);
  return response.data;
}

export async function getGradeById(gradeId) {
  const response = await axiosInstance.get(`/grades/${gradeId}`);
  return response.data;
}

export async function updateGrade(gradeId, gradeData) {
  const response = await axiosInstance.put(`/grades/act/${gradeId}`, gradeData);
  return response.data;
}

export async function deleteGrade(gradeId) {
  const response = await axiosInstance.delete(`/grades/delete/${gradeId}`);
  return response.data;
}
