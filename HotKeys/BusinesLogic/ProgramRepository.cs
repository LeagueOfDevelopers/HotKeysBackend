using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BisinessLogic
{
    class ProgramRepository : IRepository<Program>
    {
        private Dictionary<Guid,Program> programs;
        public Program Get(Guid id)
        {
            return programs[id];
        }

        public void Set(Guid id, Program item)
        {
            programs.Add(id, item);
        }
    }
}
