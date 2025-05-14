#define rand	pan_rand
#define pthread_equal(a,b)	((a)==(b))
#if defined(HAS_CODE) && defined(VERBOSE)
	#ifdef BFS_PAR
		bfs_printf("Pr: %d Tr: %d\n", II, t->forw);
	#else
		cpu_printf("Pr: %d Tr: %d\n", II, t->forw);
	#endif
#endif
	switch (t->forw) {
	default: Uerror("bad forward move");
	case 0:	/* if without executable clauses */
		continue;
	case 1: /* generic 'goto' or 'skip' */
		IfNotBlocked
		_m = 3; goto P999;
	case 2: /* generic 'else' */
		IfNotBlocked
		if (trpt->o_pm&1) continue;
		_m = 3; goto P999;

		 /* PROC :init: */
	case 3: // STATE 1 - simulation.pml:29 - [sem!0] (0:0:0 - 1)
		IfNotBlocked
		reached[1][1] = 1;
		if (q_full(now.sem))
			continue;
#ifdef HAS_CODE
		if (readtrail && gui) {
			char simtmp[64];
			sprintf(simvals, "%d!", now.sem);
		sprintf(simtmp, "%d", 0); strcat(simvals, simtmp);		}
#endif
		
		qsend(now.sem, 0, 0, 1);
		_m = 2; goto P999; /* 0 */
	case 4: // STATE 2 - simulation.pml:32 - [(run Process(1))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][2] = 1;
		if (!(addproc(II, 1, 0, 1)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 5: // STATE 3 - simulation.pml:33 - [(run Process(2))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][3] = 1;
		if (!(addproc(II, 1, 0, 2)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 6: // STATE 4 - simulation.pml:34 - [(run Process(3))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][4] = 1;
		if (!(addproc(II, 1, 0, 3)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 7: // STATE 6 - simulation.pml:36 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[1][6] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */

		 /* PROC Process */
	case 8: // STATE 1 - simulation.pml:10 - [sem?_] (0:0:1 - 1)
		reached[0][1] = 1;
		if (q_len(now.sem) == 0) continue;

		XX=1;
		(trpt+1)->bup.oval = qrecv(now.sem, XX-1, 0, 0);
		;
		qrecv(now.sem, XX-1, 0, 1);
		
#ifdef HAS_CODE
		if (readtrail && gui) {
			char simtmp[32];
			sprintf(simvals, "%d?", now.sem);
		sprintf(simtmp, "%d", ((int)_)); strcat(simvals, simtmp);		}
#endif
		;
		_m = 4; goto P999; /* 0 */
	case 9: // STATE 2 - simulation.pml:13 - [assert(!(critical))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][2] = 1;
		spin_assert( !(((int)now.critical)), " !(critical)", II, tt, t);
		_m = 3; goto P999; /* 0 */
	case 10: // STATE 3 - simulation.pml:14 - [critical = 1] (0:0:1 - 1)
		IfNotBlocked
		reached[0][3] = 1;
		(trpt+1)->bup.oval = ((int)now.critical);
		now.critical = 1;
#ifdef VAR_RANGES
		logval("critical", ((int)now.critical));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 11: // STATE 4 - simulation.pml:15 - [printf('Proceso %d ENTRANDO a sección crítica\\n',id)] (0:0:0 - 1)
		IfNotBlocked
		reached[0][4] = 1;
		Printf("Proceso %d ENTRANDO a sección crítica\n", ((int)((P0 *)_this)->id));
		_m = 3; goto P999; /* 0 */
	case 12: // STATE 5 - simulation.pml:18 - [critical = 0] (0:0:1 - 1)
		IfNotBlocked
		reached[0][5] = 1;
		(trpt+1)->bup.oval = ((int)now.critical);
		now.critical = 0;
#ifdef VAR_RANGES
		logval("critical", ((int)now.critical));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 13: // STATE 6 - simulation.pml:19 - [printf('Proceso %d SALIENDO de sección crítica\\n',id)] (0:0:0 - 1)
		IfNotBlocked
		reached[0][6] = 1;
		Printf("Proceso %d SALIENDO de sección crítica\n", ((int)((P0 *)_this)->id));
		_m = 3; goto P999; /* 0 */
	case 14: // STATE 7 - simulation.pml:21 - [sem!0] (0:0:0 - 1)
		IfNotBlocked
		reached[0][7] = 1;
		if (q_full(now.sem))
			continue;
#ifdef HAS_CODE
		if (readtrail && gui) {
			char simtmp[64];
			sprintf(simvals, "%d!", now.sem);
		sprintf(simtmp, "%d", 0); strcat(simvals, simtmp);		}
#endif
		
		qsend(now.sem, 0, 0, 1);
		_m = 2; goto P999; /* 0 */
	case 15: // STATE 11 - simulation.pml:25 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[0][11] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */
	case  _T5:	/* np_ */
		if (!((!(trpt->o_pm&4) && !(trpt->tau&128))))
			continue;
		/* else fall through */
	case  _T2:	/* true */
		_m = 3; goto P999;
#undef rand
	}

