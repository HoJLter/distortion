import { ComponentFixture, TestBed } from '@angular/core/testing';
import { JoinServer } from './join-server';

describe('JoinServer', () => {
    let component: JoinServer;
    let fixture: ComponentFixture<JoinServer>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [JoinServer],
        }).compileComponents();

        fixture = TestBed.createComponent(JoinServer);
        component = fixture.componentInstance;
        await fixture.whenStable();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
