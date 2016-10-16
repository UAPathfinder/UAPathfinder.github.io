import { Component, Input, Output, EventEmitter } from '@angular/core';

// A button with rounded edges. Used for the day picker.
@Component({
    selector: 'round-button',
	templateUrl: 'app/round-button.component.html',
	styleUrls: ['app/round-button.component.css'],
})
export class RoundButtonComponent {
	@Input() name: string;

	@Input() selected: boolean = false;
	@Output() selectedChange = new EventEmitter<boolean>();

	toggle() {
		this.selected = !this.selected;
		this.selectedChange.emit(this.selected);
	}
}

