import { Component } from '@angular/core';
import { UserService, Role, Permission } from '../../services/user.service';
import { first } from 'rxjs';
import { NgFor, NgIf } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IonicModule } from '@ionic/angular';
import { SidebarPageComponent } from '../../../../shared/stdlib/components/sidebar-page/sidebar-page.component';
import { ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';

@Component({
	selector: 'app-roles',
	templateUrl: './roles.component.html',
	styleUrls: ['./roles.component.css'],
	standalone: true,
	imports: [SidebarPageComponent, IonicModule, FormsModule, NgFor, NgIf],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RolesComponent {
	roles: Role[] = [];
	permissions: Permission[] = [];
	selectedRole: Role | undefined;
	selectedRolePermissions: Set<string> = new Set<string>();

	roleSearchQuery: string = '';
	permissionSearchQuery: string = '';

	constructor(
		private userService: UserService,
		private cd: ChangeDetectorRef
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	async loggedInInit() {
		const rsp = await this.userService.getRoles();
		this.roles = await rsp.roles;

		const rsp2 = await this.userService.getPermissions();
		this.permissions = await rsp2.permissions;

		this.cd.markForCheck();
	}

	selectRole(role: Role) {
		this.selectedRole = role;
		this.loadRolePermissions(role);
	}

	loadRolePermissions(role: Role) {
		this.selectedRolePermissions.clear();
		if (role.permissionIds) {
			for (const id of role.permissionIds) {
				this.selectedRolePermissions.add(id);
			}
		}
	}

	togglePermission(permissionId: string) {
		if (this.selectedRolePermissions.has(permissionId)) {
			this.selectedRolePermissions.delete(permissionId);
		} else {
			this.selectedRolePermissions.add(permissionId);
		}
	}

	async savePermissions() {
		if (this.selectedRole) {
			const permissionIds = [...this.selectedRolePermissions];
			await this.userService.setRolePermissions(
				this.selectedRole.id as string,
				permissionIds
			);
			// Optionally, you can provide some user feedback here, such as a success message
		}
	}

	filteredRoles() {
		return this.roles.filter((role) =>
			role.name?.toLowerCase().includes(this.roleSearchQuery.toLowerCase())
		);
	}

	filteredPermissions() {
		return this.permissions.filter((permission) =>
			permission.name
				.toLowerCase()
				.includes(this.permissionSearchQuery.toLowerCase())
		);
	}
}
