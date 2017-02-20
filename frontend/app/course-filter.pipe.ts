import { PipeTransform, Pipe } from '@angular/core';

import { Course } from './models/course';

@Pipe({
    name: 'courseFilter'
})
export class CourseFilterPipe implements PipeTransform {

    transform(value: Array<Course>, filterBy: string): Array<Course> {
        filterBy = filterBy ? filterBy.toLocaleLowerCase() : null;
        return filterBy ? value.filter((course: Course
        ) =>
            course.Title.String.toLocaleLowerCase().indexOf(filterBy) !== -1) : value;
    }
}
