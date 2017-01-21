import { Table, Column, PrimaryGeneratedColumn } from "typeorm";

/*
CREATE TABLE departments (
	id          INTEGER PRIMARY KEY AUTOINCREMENT,

	-- A short identifier
	identifier  TEXT NOCASE,

	-- A long identifier
	title       TEXT,

	description TEXT,

	UNIQUE(identifier)
);
*/

@Table("departments")
export class Department {
	@PrimaryGeneratedColumn()
	id: number;

	@Column()
	identifier: string;

	@Column()
	title: string;

	@Column()
	description: string;
}
