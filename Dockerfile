FROM golang:1.7
ADD CourseCorrect-Student /
ADD frontend/ frontend/
CMD ["/CourseCorrect-Student"]


