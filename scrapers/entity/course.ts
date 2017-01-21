import { Table, Column, PrimaryGeneratedColumn } from "typeorm";

@Table("courses")
export class Course {
	@PrimaryGeneratedColumn()
	id: number;

    @Column()
	department: string;

	@Column()
	identifier: string;

	@Column()
	title: string;

	@Column()
	description: string;

	@Column()
	units: number;
}

