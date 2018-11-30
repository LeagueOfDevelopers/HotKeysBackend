using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BisinessLogic
{
    class ProgramRepository : IRepository<Program_>
    {
        private Dictionary<Guid,Program_> programs;
        public Program_ Get(Guid id)
        {
            return programs[id];
        }

        public void Set(Guid id, Program_ item)
        {
            programs.Add(id, item);
        }
    }
}
