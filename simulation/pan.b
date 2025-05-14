	switch (t->back) {
	default: Uerror("bad return move");
	case  0: goto R999; /* nothing to undo */

		 /* PROC :init: */

	case 3: // STATE 1
		;
		_m = unsend(now.sem);
		;
		goto R999;

	case 4: // STATE 2
		;
		;
		delproc(0, now._nr_pr-1);
		;
		goto R999;

	case 5: // STATE 3
		;
		;
		delproc(0, now._nr_pr-1);
		;
		goto R999;

	case 6: // STATE 4
		;
		;
		delproc(0, now._nr_pr-1);
		;
		goto R999;

	case 7: // STATE 6
		;
		p_restor(II);
		;
		;
		goto R999;

		 /* PROC Process */

	case 8: // STATE 1
		;
		XX = 1;
		unrecv(now.sem, XX-1, 0, trpt->bup.oval, 1);
		;
		;
		goto R999;
;
		;
		
	case 10: // STATE 3
		;
		now.critical = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 12: // STATE 5
		;
		now.critical = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 14: // STATE 7
		;
		_m = unsend(now.sem);
		;
		goto R999;

	case 15: // STATE 11
		;
		p_restor(II);
		;
		;
		goto R999;
	}

