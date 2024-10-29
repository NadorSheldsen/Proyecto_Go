// src/services/subjectsService.js
import axiosInstance from '../axiosConfig';

export async function getAllSubjects() {
  const response = await axiosInstance.get('/subjects');
  return response.data;
}

export async function createSubject(subjectData) {
  const response = await axiosInstance.post('/escuela/subjects', subjectData);
  return response.data;
}

export async function getSubjectById(subjectId) {
  const response = await axiosInstance.get(`/subjects/${subjectId}`);
  return response.data;
}

export async function updateSubject(subjectId, subjectData) {
  const response = await axiosInstance.put(`/escuela/subjects/${subjectId}`, subjectData);
  return response.data;
}

export async function deleteSubject(subjectId) {
  const response = await axiosInstance.delete(`/escuela/subjects/${subjectId}`);
  return response.data;
}
