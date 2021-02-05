#pragma once

#include <memory>

#include "src/shared/types/types.h"
#include "src/stirling/bpf_tools/bcc_wrapper.h"
#include "src/stirling/core/source_connector.h"
#include "src/stirling/core/types.h"
#include "src/stirling/source_connectors/perf_profiler/bcc_bpf_intf/stack_event.h"
#include "src/stirling/source_connectors/perf_profiler/stack_traces_table.h"

namespace pl {
namespace stirling {

class PerfProfileConnector : public SourceConnector, public bpf_tools::BCCWrapper {
 public:
  static constexpr auto kTables = MakeArray(kStackTraceTable);

  static std::unique_ptr<SourceConnector> Create(std::string_view name) {
    return std::unique_ptr<SourceConnector>(new PerfProfileConnector(name));
  }

  Status InitImpl() override;
  Status StopImpl() override;
  void TransferDataImpl(ConnectorContext* ctx, uint32_t table_num, DataTable* data_table) override;

 private:
  static constexpr auto kProbeSpecs = MakeArray<bpf_tools::KProbeSpec>({
      {"sample_call_stack", bpf_tools::BPFProbeAttachType::kEntry, "syscall__probe_entry_connect"},
  });

  explicit PerfProfileConnector(std::string_view source_name);

  std::unique_ptr<ebpf::BPFStackTable> stacks_;
  std::unique_ptr<ebpf::BPFHashTable<stack_trace_key_t, uint64_t> > counts_;
};

}  // namespace stirling
}  // namespace pl
