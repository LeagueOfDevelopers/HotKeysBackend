using HotKeysLibrary.Entities;
using System;
using System.Collections.Generic;

namespace HotKeysLibrary.Infrastructure
{
    public class ProgramRepository : IRepository<Program>
    {
        private readonly Dictionary<Guid, Program> _programs;

        public ProgramRepository(Dictionary<Guid, Program> programs)
        {
            _programs = programs;
        }

        public Program Get(Guid id)
        {
            return _programs[id];
        }

        public void Add(Guid id, Program program)
        {
            if (program == null)
                throw new NullReferenceException();
            _programs.Add(id, program);
        }
    }
}
