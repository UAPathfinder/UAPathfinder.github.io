import { Table, Column, PrimaryGeneratedColumn } from "typeorm";

@Table("classes")
export class Class {
	@PrimaryGeneratedColumn()
	id: number;

    @Column()
	identifier: string;

	@Column()
	course: string;

	@Column()
	location: string;

	@Column()
	capicity: number;

	@Column()
	registered: number;

	@Column()
	professor: string;

	@Column()
	start_time: number;

	@Column()
	end_time: number;

	@Column()
	start_date: number;

	@Column()
	end_date: number;

	@Column()
	sunday: boolean;

	@Column()
	monday: boolean;

	@Column()
	tuesday: boolean;

	@Column()
	wednesday: boolean;

	@Column()
	thursday: boolean;

	@Column()
	friday: boolean;

	@Column()
	saturday: boolean;
}

