import {
	Component,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
} from '@angular/core';
import { NavController, IonicModule } from '@ionic/angular';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { map } from 'rxjs/operators';
import { NgStyle, NgIf } from '@angular/common';

type appGroup = 'ai-group' | 'users-group' | '';

@Component({
	selector: 'app-icon-menu',
	standalone: true,
	imports: [IonicModule, NgStyle, RouterLink, NgIf],
	templateUrl: './icon-menu.component.html',
	styleUrl: './icon-menu.component.scss',
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class IconMenuComponent {
	currentPath = '';

	constructor(
		public navCtrl: NavController,
		private activatedRoute: ActivatedRoute,
		private cd: ChangeDetectorRef
	) {}

	ngOnInit() {
		this.activatedRoute.url
			.pipe(map((segments) => segments.join('/')))
			.subscribe((path) => {
				this.currentPath = path;
				this.cd.markForCheck();
			});
	}

	group(): appGroup {
		if (
			this.currentPath === 'startup' ||
			this.currentPath === 'chat' ||
			this.currentPath === 'model-explorer'
		) {
			return 'ai-group';
		}

		if (
			this.currentPath === 'users' ||
			this.currentPath === 'add-user' ||
			this.currentPath === 'roles'
		) {
			return 'users-group';
		}

		return '';
	}

	rout = {
		activeEntry: '',
	};
}
