import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { SortablejsModule } from 'angular-sortablejs';

import { AppComponent }  from './app.component';
import { RoundButtonComponent }  from './round-button.component';

import { CourseService } from './course.service';
import { TimePipe } from './time.pipe';
import { CourseFilterPipe } from './course-filter.pipe';

@NgModule({
  imports: [ BrowserModule, FormsModule, HttpModule, SortablejsModule ],
  declarations: [ AppComponent, RoundButtonComponent, TimePipe, CourseFilterPipe ],
  bootstrap: [ AppComponent ],
  providers: [ CourseService ],
})
export class AppModule { }
