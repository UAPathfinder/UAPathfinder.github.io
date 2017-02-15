FROM golang:1.7
ADD CourseCorrect-Student /
ADD data/test data/
CMD ["/CourseCorrect-Student"]

