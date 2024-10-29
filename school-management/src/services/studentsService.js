// src/services/studentsService.js
import axiosInstance from '../axiosConfig';

export async function getAllStudents() {
  const response = await axiosInstance.get('/students');
  return response.data;
}

export async function createStudent(studentData) {
  const response = await axiosInstance.post('/bd/students', studentData);
  return response.data;
}

export async function getStudentById(studentId) {
  const response = await axiosInstance.get(`/students/${studentId}`);
  return response.data;
}

export async function updateStudent(studentId, studentData) {
  const response = await axiosInstance.put(`/bd/students/update/${studentId}`, studentData);
  return response.data;
}

export async function deleteStudent(studentId) {
  const response = await axiosInstance.post(`/bd/students/delete/${studentId}`);
  return response.data;
}
